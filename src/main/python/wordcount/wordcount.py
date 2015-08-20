def wordcount(message):
    counts = {}
    print message
    words = message.split(" ")
    for word in words:
        if counts.has_key(word):
            counts[word] += 1
        else:
            counts[word] = 1
    return counts

if __name__ == "__main__":
    print "{0}".format(wordcount("This exercise from the golang tour This tour is very very very ... "))
