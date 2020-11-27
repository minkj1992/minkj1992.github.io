from django.http import HttpResponse, JsonResponse
from django.views.decorators.csrf import csrf_exempt
from rest_framework.parsers import JSONParser
from snippets.models import Snippet
from snippets.serializers import SnippetSerializer


# exempt -> exemption 면제
@csrf_exempt
def snippet_list(request):
    """List all code snippets or crete a new snippet."""
    if request.method == 'GET':
        snippets = Snippet.objects.all()
        serializer = SnippetSerializer(snippets, many=True)
        return JsonResponse(serializer.data, safe=False)
    
    elif request.method == 'POST':
        data = JSONParser.parse(request)
        serializer = SnippetSerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            return JsonResponse(serializer.data, status=201) 
        return JsonResponse(serializer.data, status=400) #  Bad Request

@csrf_exempt
def snippet_detail(request, pk):
    """Retrieve, Update, Delete a code-snippet."""
    try:
        snippet = Snippet.objects.get(pk=pk)
    except Snippet.DoesNotExist:
        return HttpResponse(status=404) # Not Found

    if request.method == 'GET':
        serializer = SnippetSerializer(snippet)
        return JsonResponse(serializer.data)

    elif request.method == 'PUT':
        data = JSONParser().parse(request) # Binary -> JSON
        serializer = SnippetSerializer(snippet, data=data)
        if serializer.is_valid():
            serializer.save()
            return JsonResponse(serializer.data)
        return JsonResponse(serializer.errors, status=400)

    elif request.method == 'DELETE':
        snippet.delete()
        return HttpResponse(status=204) #No Content



def create_serializer():
    """This is for test code to generate snippet and serializer"""
    for snippet in SnippetSerializer.objects.all():
        serializer = SnippetSerializer(snippet)
        content = JSONRenderer().render(serializer.data)
        stream = io.BytesIO(content)
        data = JSONParser().parse(stream)
        serializer = SnippetSerializer(data=data)
        serializer.is_valid()
        serializer.save()